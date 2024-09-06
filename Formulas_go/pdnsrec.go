package main

// Pdnsrec - Non-authoritative/recursing DNS server
// Homepage: https://www.powerdns.com/powerdns-recursor

import (
	"fmt"
	
	"os/exec"
)

func installPdnsrec() {
	// Método 1: Descargar y extraer .tar.gz
	pdnsrec_tar_url := "https://downloads.powerdns.com/releases/pdns-recursor-5.1.1.tar.bz2"
	pdnsrec_cmd_tar := exec.Command("curl", "-L", pdnsrec_tar_url, "-o", "package.tar.gz")
	err := pdnsrec_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pdnsrec_zip_url := "https://downloads.powerdns.com/releases/pdns-recursor-5.1.1.tar.bz2"
	pdnsrec_cmd_zip := exec.Command("curl", "-L", pdnsrec_zip_url, "-o", "package.zip")
	err = pdnsrec_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pdnsrec_bin_url := "https://downloads.powerdns.com/releases/pdns-recursor-5.1.1.tar.bz2"
	pdnsrec_cmd_bin := exec.Command("curl", "-L", pdnsrec_bin_url, "-o", "binary.bin")
	err = pdnsrec_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pdnsrec_src_url := "https://downloads.powerdns.com/releases/pdns-recursor-5.1.1.tar.bz2"
	pdnsrec_cmd_src := exec.Command("curl", "-L", pdnsrec_src_url, "-o", "source.tar.gz")
	err = pdnsrec_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pdnsrec_cmd_direct := exec.Command("./binary")
	err = pdnsrec_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: lua")
exec.Command("latte", "install", "lua")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
}
