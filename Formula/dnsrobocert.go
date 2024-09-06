package main

// Dnsrobocert - Manage Let's Encrypt SSL certificates based on DNS challenges
// Homepage: https://dnsrobocert.readthedocs.io/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installDnsrobocert() {
	// Método 1: Descargar y extraer .tar.gz
	dnsrobocert_tar_url := "https://files.pythonhosted.org/packages/85/2a/bf24b7888b5be61a7ffcfb75d426f05f9a11f3a3262a12887d546943b70f/dnsrobocert-3.25.0.tar.gz"
	dnsrobocert_cmd_tar := exec.Command("curl", "-L", dnsrobocert_tar_url, "-o", "package.tar.gz")
	err := dnsrobocert_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dnsrobocert_zip_url := "https://files.pythonhosted.org/packages/85/2a/bf24b7888b5be61a7ffcfb75d426f05f9a11f3a3262a12887d546943b70f/dnsrobocert-3.25.0.zip"
	dnsrobocert_cmd_zip := exec.Command("curl", "-L", dnsrobocert_zip_url, "-o", "package.zip")
	err = dnsrobocert_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dnsrobocert_bin_url := "https://files.pythonhosted.org/packages/85/2a/bf24b7888b5be61a7ffcfb75d426f05f9a11f3a3262a12887d546943b70f/dnsrobocert-3.25.0.bin"
	dnsrobocert_cmd_bin := exec.Command("curl", "-L", dnsrobocert_bin_url, "-o", "binary.bin")
	err = dnsrobocert_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dnsrobocert_src_url := "https://files.pythonhosted.org/packages/85/2a/bf24b7888b5be61a7ffcfb75d426f05f9a11f3a3262a12887d546943b70f/dnsrobocert-3.25.0.src.tar.gz"
	dnsrobocert_cmd_src := exec.Command("curl", "-L", dnsrobocert_src_url, "-o", "source.tar.gz")
	err = dnsrobocert_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dnsrobocert_cmd_direct := exec.Command("./binary")
	err = dnsrobocert_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
