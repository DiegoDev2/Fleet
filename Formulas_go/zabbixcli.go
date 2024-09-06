package main

// ZabbixCli - CLI tool for interacting with Zabbix monitoring system
// Homepage: https://unioslo.github.io/zabbix-cli/

import (
	"fmt"
	
	"os/exec"
)

func installZabbixCli() {
	// Método 1: Descargar y extraer .tar.gz
	zabbixcli_tar_url := "https://github.com/unioslo/zabbix-cli/archive/refs/tags/3.0.1.tar.gz"
	zabbixcli_cmd_tar := exec.Command("curl", "-L", zabbixcli_tar_url, "-o", "package.tar.gz")
	err := zabbixcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zabbixcli_zip_url := "https://github.com/unioslo/zabbix-cli/archive/refs/tags/3.0.1.zip"
	zabbixcli_cmd_zip := exec.Command("curl", "-L", zabbixcli_zip_url, "-o", "package.zip")
	err = zabbixcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zabbixcli_bin_url := "https://github.com/unioslo/zabbix-cli/archive/refs/tags/3.0.1.bin"
	zabbixcli_cmd_bin := exec.Command("curl", "-L", zabbixcli_bin_url, "-o", "binary.bin")
	err = zabbixcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zabbixcli_src_url := "https://github.com/unioslo/zabbix-cli/archive/refs/tags/3.0.1.src.tar.gz"
	zabbixcli_cmd_src := exec.Command("curl", "-L", zabbixcli_src_url, "-o", "source.tar.gz")
	err = zabbixcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zabbixcli_cmd_direct := exec.Command("./binary")
	err = zabbixcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
