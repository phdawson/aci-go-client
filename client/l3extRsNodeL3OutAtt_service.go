package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateL3extRsNodeL3OutAtt(tDn string, logical_node_profile string, l3_outside string, tenant string, description string, l3extRsNodeL3OutAttattr models.L3extRsNodeL3OutAttAttributes) (*models.L3extRsNodeL3OutAtt, error) {
	rn := fmt.Sprintf("rsnodeL3OutAtt-[%s]", tDn)
	parentDn := fmt.Sprintf("uni/tn-%s/out-%s/lnodep-%s", tenant, l3_outside, logical_node_profile)
	l3extRsNodeL3OutAtt := models.NewL3extRsNodeL3OutAtt(rn, parentDn, description, l3extRsNodeL3OutAttattr)
	err := sm.Save(l3extRsNodeL3OutAtt)
	return l3extRsNodeL3OutAtt, err
}

func (sm *ServiceManager) ReadL3extRsNodeL3OutAtt(tDn string, logical_node_profile string, l3_outside string, tenant string) (*models.L3extRsNodeL3OutAtt, error) {
	dn := fmt.Sprintf("uni/tn-%s/out-%s/lnodep-%s/rsnodeL3OutAtt-[%s]", tenant, l3_outside, logical_node_profile, tDn)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	l3extRsNodeL3OutAtt := models.L3extRsNodeL3OutAttFromContainer(cont)
	return l3extRsNodeL3OutAtt, nil
}

func (sm *ServiceManager) DeleteL3extRsNodeL3OutAtt(tDn string, logical_node_profile string, l3_outside string, tenant string) error {
	dn := fmt.Sprintf("uni/tn-%s/out-%s/lnodep-%s/rsnodeL3OutAtt-[%s]", tenant, l3_outside, logical_node_profile, tDn)
	return sm.DeleteByDn(dn, models.L3extrsnodel3outattClassName)
}

func (sm *ServiceManager) UpdateL3extRsNodeL3OutAtt(tDn string, logical_node_profile string, l3_outside string, tenant string, description string, l3extRsNodeL3OutAttattr models.L3extRsNodeL3OutAttAttributes) (*models.L3extRsNodeL3OutAtt, error) {
	rn := fmt.Sprintf("rsnodeL3OutAtt-[%s]", tDn)
	parentDn := fmt.Sprintf("uni/tn-%s/out-%s/lnodep-%s", tenant, l3_outside, logical_node_profile)
	l3extRsNodeL3OutAtt := models.NewL3extRsNodeL3OutAtt(rn, parentDn, description, l3extRsNodeL3OutAttattr)

	l3extRsNodeL3OutAtt.Status = "modified"
	err := sm.Save(l3extRsNodeL3OutAtt)
	return l3extRsNodeL3OutAtt, err

}

func (sm *ServiceManager) ListL3extRsNodeL3OutAtt(logical_node_profile string, l3_outside string, tenant string) ([]*models.L3extRsNodeL3OutAtt, error) {

	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/out-%s/lnodep-%s/l3extRsNodeL3OutAtt.json", baseurlStr, tenant, l3_outside, logical_node_profile)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.L3extRsNodeL3OutAttListFromContainer(cont)

	return list, err
}
