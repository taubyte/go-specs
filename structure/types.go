package structureSpec

type Structure interface {
	*App |
		*Database |
		*Domain |
		*Function |
		*Library |
		*Messaging |
		*Service |
		*SmartOp |
		*Storage |
		*Website |

		// Added for usage outside of defined resources
		*Simple

	GetName() string
	SetId(string)
}
