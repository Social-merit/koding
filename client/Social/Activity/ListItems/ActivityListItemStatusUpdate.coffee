class StatusActivityItemView extends ActivityItemChild
  constructor:(options = {}, data={})->
    options.cssClass or= "activity-item status"
    options.tooltip  or=
      title            : "Status Update"
      selector         : "span.type-icon"
      offset           :
        top            : 3
        left           : -5

    super options, data

    embedOptions  =
      hasDropdown : no
      delegate    : this

    if data.link?
      @embedBox = new EmbedBox embedOptions, data.link
      @setClass "two-columns"  if @twoColumns
    else
      @embedBox = new KDView

    @timeAgoView = new KDTimeAgoView {}, @getData().meta.createdAt

  getTokenMap: (tokens) ->
    return  unless tokens
    map = {}
    tokens.forEach (token) -> map[token.getId()] = token
    return  map

  expandTokens: (str = "") ->
    return  str unless tokenMatches = str.match /\|.+?\|/g

    data = @getData()
    tagMap = @getTokenMap data.tags  if data.tags

    viewParams = []
    for tokenString in tokenMatches
      [prefix, constructorName, id] = @decodeToken tokenString

      switch prefix
        when "#" then token = tagMap?[id]
        else continue

      continue  unless token

      domId     = @utils.getUniqueId()
      itemClass = tokenClassMap[prefix]
      tokenView = new TokenView {domId, itemClass}, token
      tokenView.emit "viewAppended"
      str = str.replace tokenString, tokenView.getElement().outerHTML
      tokenView.destroy()

      viewParams.push {options: {domId, itemClass}, data: token}

    @utils.defer ->
      for {options, data} in viewParams
        new TokenView options, data

    return  str

  decodeToken: (str) ->
    return  match[1].split /:/g  if match = str.match /^\|(.+)\|$/

  formatContent: (str = "")->
    str = @utils.applyMarkdown str
    str = @expandTokens str
    return  str

  viewAppended:->
    return if @getData().constructor is KD.remote.api.CStatusActivity
    super
    @setTemplate @pistachio()
    @template.update()

    @utils.defer =>
      predicate = @getData().link?.link_url? and @getData().link.link_url isnt ''
      if predicate
      then @embedBox.show()
      else @embedBox.hide()

  pistachio:->
    """
      {{> @avatar}}
      {{> @settingsButton}}
      {{> @author}}
      <p class="status-body">{{@formatContent #(body)}}</p>
      {{> @embedBox}}
      <footer>
        {{> @actionLinks}}
        {{> @timeAgoView}}
      </footer>
      {{> @commentBox}}
    """

  tokenClassMap =
    "#"         : TagLinkView
