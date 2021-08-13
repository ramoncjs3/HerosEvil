package app

import "errors"

func ChooseServer(servercode string) error {
	switch servercode {
	case "g1":
		LoginStruct.QuickLoginUrl = `http://122.152.210.142:8002`
		LoginStruct.KillNpcUrl = `http://122.152.210.142:9002`
	case "g2":
		LoginStruct.QuickLoginUrl = `http://122.51.2.7:8002`
		LoginStruct.KillNpcUrl = `http://122.51.2.7:9002`

	case "g3":
		LoginStruct.QuickLoginUrl = `http://115.159.83.24:8002`
		LoginStruct.KillNpcUrl = `http://115.159.83.24:9002`

	case "h1":
		LoginStruct.QuickLoginUrl = `http://122.51.2.7:8003`
		LoginStruct.KillNpcUrl = `http://122.51.2.7:9003`

	case "h2":
		LoginStruct.QuickLoginUrl = `http://122.152.207.125:8003`
		LoginStruct.KillNpcUrl = `http://122.152.207.125:9003`

	case "h3":
		LoginStruct.QuickLoginUrl = `http://115.159.83.24:8003`
		LoginStruct.KillNpcUrl = `http://115.159.83.24:9003`

	case "h4":
		LoginStruct.QuickLoginUrl = `http://111.231.91.252:8004`
		LoginStruct.KillNpcUrl = `http://111.231.91.252:9004`

	case "h5":
		LoginStruct.QuickLoginUrl = `http://122.51.2.7:8004`
		LoginStruct.KillNpcUrl = `http://122.51.2.7:9004`

	case "h6":
		LoginStruct.QuickLoginUrl = `http://122.152.210.142:8004`
		LoginStruct.KillNpcUrl = `http://122.152.210.142:9004`

	case "h7":
		LoginStruct.QuickLoginUrl = `http://106.54.75.58:8002`
		LoginStruct.KillNpcUrl = `http://106.54.75.58:9002`

	case "h8":
		LoginStruct.QuickLoginUrl = `http://175.24.32.165:8002`
		LoginStruct.KillNpcUrl = `http://175.24.32.165:9002`

	case "h9":
		LoginStruct.QuickLoginUrl = `http://175.24.32.165:8003`
		LoginStruct.KillNpcUrl = `http://175.24.32.165:9003`

	case "h10":
		LoginStruct.QuickLoginUrl = `http://101.34.70.93:8002`
		LoginStruct.KillNpcUrl = `http://101.34.70.93:9002`
	default:
		return errors.New("区服代码错误,通过--help获取帮助.")
	}
	return nil
}
