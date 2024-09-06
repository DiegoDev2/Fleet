package main

// Zabbix - Availability and monitoring solution
// Homepage: https://www.zabbix.com/

import (
	"fmt"
	
	"os/exec"
)

func installZabbix() {
	// Método 1: Descargar y extraer .tar.gz
	zabbix_tar_url := "https://cdn.zabbix.com/zabbix/sources/stable/7.0/zabbix-7.0.3.tar.gz"
	zabbix_cmd_tar := exec.Command("curl", "-L", zabbix_tar_url, "-o", "package.tar.gz")
	err := zabbix_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zabbix_zip_url := "https://cdn.zabbix.com/zabbix/sources/stable/7.0/zabbix-7.0.3.zip"
	zabbix_cmd_zip := exec.Command("curl", "-L", zabbix_zip_url, "-o", "package.zip")
	err = zabbix_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zabbix_bin_url := "https://cdn.zabbix.com/zabbix/sources/stable/7.0/zabbix-7.0.3.bin"
	zabbix_cmd_bin := exec.Command("curl", "-L", zabbix_bin_url, "-o", "binary.bin")
	err = zabbix_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zabbix_src_url := "https://cdn.zabbix.com/zabbix/sources/stable/7.0/zabbix-7.0.3.src.tar.gz"
	zabbix_cmd_src := exec.Command("curl", "-L", zabbix_src_url, "-o", "source.tar.gz")
	err = zabbix_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zabbix_cmd_direct := exec.Command("./binary")
	err = zabbix_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
}
