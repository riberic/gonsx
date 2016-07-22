package securitypolicy


// AddSecurityGroupBinding - Adds security group to list of SecurityGroupBinding if it doesn't exists.
func (sp *SecurityPolicy) AddSecurityGroupBinding(objectID string){
	for _, secGroup := range sp.SecurityGroupBinding {
		if secGroup.ObjectID == objectID {
			return
		}
	}
	// if we reached here that means we couldn't find one, and let's add the sec group.
	sp.SecurityGroupBinding = append(sp.SecurityGroupBinding, SecurityGroup{ObjectID: objectID})
	return
}


// RemoveSecurityGroupBinding - Adds security group to list of SecurityGroupBinding if it doesn't exists.
func (sp *SecurityPolicy) RemoveSecurityGroupBinding(objectID string){
	for idx, secGroup := range sp.SecurityGroupBinding {
		if secGroup.ObjectID == objectID {
			sp.SecurityGroupBinding = append(sp.SecurityGroupBinding[:idx], sp.SecurityGroupBinding[idx+1:]...)
			return
		}
	}
	return
}
