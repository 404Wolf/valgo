package valgo

func (v *ExtendedVal) ExtendedValToBasicVal() BasicVal {
	return *NewBasicVal(
		v.Name,
		v.Id,
		v.Version,
		v.Code,
		v.Public,
		v.CreatedAt,
		v.Privacy,
		v.Type,
		v.Url,
		v.Links,
		v.Author,
	)
}
