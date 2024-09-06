package main

// Suricata - Network IDS, IPS, and security monitoring engine
// Homepage: https://suricata.io

import (
	"fmt"
	
	"os/exec"
)

func installSuricata() {
	// Método 1: Descargar y extraer .tar.gz
	suricata_tar_url := "https://www.openinfosecfoundation.org/download/suricata-7.0.6.tar.gz"
	suricata_cmd_tar := exec.Command("curl", "-L", suricata_tar_url, "-o", "package.tar.gz")
	err := suricata_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	suricata_zip_url := "https://www.openinfosecfoundation.org/download/suricata-7.0.6.zip"
	suricata_cmd_zip := exec.Command("curl", "-L", suricata_zip_url, "-o", "package.zip")
	err = suricata_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	suricata_bin_url := "https://www.openinfosecfoundation.org/download/suricata-7.0.6.bin"
	suricata_cmd_bin := exec.Command("curl", "-L", suricata_bin_url, "-o", "binary.bin")
	err = suricata_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	suricata_src_url := "https://www.openinfosecfoundation.org/download/suricata-7.0.6.src.tar.gz"
	suricata_cmd_src := exec.Command("curl", "-L", suricata_src_url, "-o", "source.tar.gz")
	err = suricata_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	suricata_cmd_direct := exec.Command("./binary")
	err = suricata_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: jansson")
exec.Command("latte", "install", "jansson")
	fmt.Println("Instalando dependencia: libmagic")
exec.Command("latte", "install", "libmagic")
	fmt.Println("Instalando dependencia: libnet")
exec.Command("latte", "install", "libnet")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
