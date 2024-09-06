package main

// Asn - Organization lookup and server tool (ASN / IPv4 / IPv6 / Prefix / AS Path)
// Homepage: https://github.com/nitefood/asn

import (
	"fmt"
	
	"os/exec"
)

func installAsn() {
	// Método 1: Descargar y extraer .tar.gz
	asn_tar_url := "https://github.com/nitefood/asn/archive/refs/tags/v0.78.0.tar.gz"
	asn_cmd_tar := exec.Command("curl", "-L", asn_tar_url, "-o", "package.tar.gz")
	err := asn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	asn_zip_url := "https://github.com/nitefood/asn/archive/refs/tags/v0.78.0.zip"
	asn_cmd_zip := exec.Command("curl", "-L", asn_zip_url, "-o", "package.zip")
	err = asn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	asn_bin_url := "https://github.com/nitefood/asn/archive/refs/tags/v0.78.0.bin"
	asn_cmd_bin := exec.Command("curl", "-L", asn_bin_url, "-o", "binary.bin")
	err = asn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	asn_src_url := "https://github.com/nitefood/asn/archive/refs/tags/v0.78.0.src.tar.gz"
	asn_cmd_src := exec.Command("curl", "-L", asn_src_url, "-o", "source.tar.gz")
	err = asn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	asn_cmd_direct := exec.Command("./binary")
	err = asn_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: aha")
	exec.Command("latte", "install", "aha").Run()
	fmt.Println("Instalando dependencia: bash")
	exec.Command("latte", "install", "bash").Run()
	fmt.Println("Instalando dependencia: coreutils")
	exec.Command("latte", "install", "coreutils").Run()
	fmt.Println("Instalando dependencia: grepcidr")
	exec.Command("latte", "install", "grepcidr").Run()
	fmt.Println("Instalando dependencia: ipcalc")
	exec.Command("latte", "install", "ipcalc").Run()
	fmt.Println("Instalando dependencia: jq")
	exec.Command("latte", "install", "jq").Run()
	fmt.Println("Instalando dependencia: mtr")
	exec.Command("latte", "install", "mtr").Run()
	fmt.Println("Instalando dependencia: nmap")
	exec.Command("latte", "install", "nmap").Run()
	fmt.Println("Instalando dependencia: bind")
	exec.Command("latte", "install", "bind").Run()
}
