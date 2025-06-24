package request

func (req *RequestImpl) fetchUserAccountData() *RequestImpl {
	if !req.securityHeaders.IsAuth {
		return req
	}

	req.userAccount = UserAccount{
		// ID: ,
	}
	return req
}
