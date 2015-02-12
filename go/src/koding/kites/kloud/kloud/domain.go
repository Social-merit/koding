package kloud

import (
	"fmt"
	"koding/kites/kloud/eventer"
	"koding/kites/kloud/protocol"
	"strings"

	"github.com/koding/kite"
)

type domainArgs struct {
	DomainName string
}

type domainFunc func(*protocol.Machine, *domainArgs) (interface{}, error)

func (k *Kloud) domainHandler(r *kite.Request, fn domainFunc) (resp interface{}, err error) {
	args := &domainArgs{}
	if err := r.Args.One().Unmarshal(args); err != nil {
		return nil, err
	}

	if err := k.Domainer.Validate(args.DomainName, r.Username); err != nil {
		return nil, err
	}

	m, err := k.PrepareMachine(r)
	if err != nil {
		return nil, err
	}

	// PrepareMachine is locking for us, so unlock after we are done
	defer k.Locker.Unlock(m.Id)

	if m.IpAddress == "" {
		return nil, fmt.Errorf("ip address is not defined")
	}

	// fake eventer to avoid panics if someone tries to use the eventer
	m.Eventer = &eventer.Events{}

	k.Log.Debug("'%s' method is called with args: %+v\n", r.Method, args)

	return fn(m, args)
}

func (k *Kloud) DomainAdd(r *kite.Request) (resp interface{}, reqErr error) {
	addFunc := func(m *protocol.Machine, args *domainArgs) (interface{}, error) {
		// a non nil means the domain exists
		if _, err := k.Domainer.Get(args.DomainName); err == nil {
			return nil, fmt.Errorf("domain record already exists")
		}

		// now assign the machine ip to the given domain name
		if err := k.Domainer.Create(args.DomainName, m.IpAddress); err != nil {
			return nil, err
		}

		domain := &protocol.Domain{
			Username:  m.Username,
			MachineId: m.Id,
			Name:      args.DomainName,
		}

		if err := k.DomainStorage.Add(domain); err != nil {
			return nil, err
		}

		return true, nil
	}

	return k.domainHandler(r, addFunc)
}

func (k *Kloud) DomainRemove(r *kite.Request) (resp interface{}, reqErr error) {
	removeFunc := func(m *protocol.Machine, args *domainArgs) (interface{}, error) {
		// do not return on error because it might be already delete via unset
		k.Domainer.Delete(args.DomainName, m.IpAddress)

		if err := k.DomainStorage.Delete(args.DomainName); err != nil {
			return nil, err
		}

		return true, nil
	}

	return k.domainHandler(r, removeFunc)
}

func (k *Kloud) DomainUnset(r *kite.Request) (resp interface{}, reqErr error) {
	unsetFunc := func(m *protocol.Machine, args *domainArgs) (interface{}, error) {
		// be sure the domain does exist in storage before we delete the domain
		if _, err := k.DomainStorage.Get(args.DomainName); err != nil {
			return nil, fmt.Errorf("domain does not exists in DB")
		}

		if err := k.Domainer.Delete(args.DomainName, m.IpAddress); err != nil {
			return nil, err
		}

		// remove the machineID for the given domain. The document is still
		// there, but it'sn associated anymore with this domain.
		if err := k.DomainStorage.UpdateMachine(args.DomainName, ""); err != nil {
			return nil, err
		}

		return true, nil
	}

	return k.domainHandler(r, unsetFunc)
}

func (k *Kloud) DomainSet(r *kite.Request) (resp interface{}, reqErr error) {
	setFunc := func(m *protocol.Machine, args *domainArgs) (interface{}, error) {
		// be sure the domain does exist in storage before we create the domain
		if _, err := k.DomainStorage.Get(args.DomainName); err != nil {
			return nil, fmt.Errorf("domain does not exists in DB")
		}

		record, err := k.Domainer.Get(args.DomainName)
		if err != nil && strings.Contains(err.Error(), "no records available") {
			k.Log.Debug("[%s] setting domain '%s' to IP '%s'", m.Id, args.DomainName, m.IpAddress)
			if err := k.Domainer.Create(args.DomainName, m.IpAddress); err != nil {
				return nil, err
			}
		} else if err != nil {
			// If it's something else just return it
			return nil, err
		}

		// check for err again, otherwise we get a panic by acessing record's fields
		if err == nil && record.IP != m.IpAddress {
			k.Log.Debug("[%s] setting domain '%s' from old IP '%s' to new Ip '%s'",
				m.Id, args.DomainName, record.IP, m.IpAddress)
			if err := k.Domainer.Update(args.DomainName, record.IP, m.IpAddress); err != nil {
				fmt.Printf("err = %+v\n", err)
				return nil, err
			}
		}

		// adding the machineID for the given domain.
		if err := k.DomainStorage.UpdateMachine(args.DomainName, m.Id); err != nil {
			return nil, err
		}

		return true, nil
	}

	return k.domainHandler(r, setFunc)
}
