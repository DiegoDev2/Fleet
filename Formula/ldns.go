package main

// Ldns - DNS library written in C
// Homepage: https://nlnetlabs.nl/projects/ldns/

import (
	"fmt"
	
	"os/exec"
)

func installLdns() {
	// Método 1: Descargar y extraer .tar.gz
	ldns_tar_url := "https://nlnetlabs.nl/downloads/ldns/ldns-1.8.4.tar.gz"
	ldns_cmd_tar := exec.Command("curl", "-L", ldns_tar_url, "-o", "package.tar.gz")
	err := ldns_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ldns_zip_url := "https://nlnetlabs.nl/downloads/ldns/ldns-1.8.4.zip"
	ldns_cmd_zip := exec.Command("curl", "-L", ldns_zip_url, "-o", "package.zip")
	err = ldns_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ldns_bin_url := "https://nlnetlabs.nl/downloads/ldns/ldns-1.8.4.bin"
	ldns_cmd_bin := exec.Command("curl", "-L", ldns_bin_url, "-o", "binary.bin")
	err = ldns_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ldns_src_url := "https://nlnetlabs.nl/downloads/ldns/ldns-1.8.4.src.tar.gz"
	ldns_cmd_src := exec.Command("curl", "-L", ldns_src_url, "-o", "source.tar.gz")
	err = ldns_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ldns_cmd_direct := exec.Command("./binary")
	err = ldns_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
	fmt.Println("Instalando dependencia: swig")
	exec.Command("latte", "install", "swig").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
