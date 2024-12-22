package valgo

func (v *ExtendedVal) ToBasicVal() BasicVal {
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
