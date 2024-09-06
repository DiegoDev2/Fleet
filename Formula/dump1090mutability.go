package main

// Dump1090Mutability - ADS-B Ground Station System for RTL-SDR
// Homepage: https://packages.ubuntu.com/jammy/dump1090-mutability

import (
	"fmt"
	
	"os/exec"
)

func installDump1090Mutability() {
	// Método 1: Descargar y extraer .tar.gz
	dump1090mutability_tar_url := "http://archive.ubuntu.com/ubuntu/pool/universe/d/dump1090-mutability/dump1090-mutability_1.15~20180310.4a16df3+dfsg.orig.tar.gz"
	dump1090mutability_cmd_tar := exec.Command("curl", "-L", dump1090mutability_tar_url, "-o", "package.tar.gz")
	err := dump1090mutability_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dump1090mutability_zip_url := "http://archive.ubuntu.com/ubuntu/pool/universe/d/dump1090-mutability/dump1090-mutability_1.15~20180310.4a16df3+dfsg.orig.zip"
	dump1090mutability_cmd_zip := exec.Command("curl", "-L", dump1090mutability_zip_url, "-o", "package.zip")
	err = dump1090mutability_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dump1090mutability_bin_url := "http://archive.ubuntu.com/ubuntu/pool/universe/d/dump1090-mutability/dump1090-mutability_1.15~20180310.4a16df3+dfsg.orig.bin"
	dump1090mutability_cmd_bin := exec.Command("curl", "-L", dump1090mutability_bin_url, "-o", "binary.bin")
	err = dump1090mutability_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dump1090mutability_src_url := "http://archive.ubuntu.com/ubuntu/pool/universe/d/dump1090-mutability/dump1090-mutability_1.15~20180310.4a16df3+dfsg.orig.src.tar.gz"
	dump1090mutability_cmd_src := exec.Command("curl", "-L", dump1090mutability_src_url, "-o", "source.tar.gz")
	err = dump1090mutability_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dump1090mutability_cmd_direct := exec.Command("./binary")
	err = dump1090mutability_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: librtlsdr")
	exec.Command("latte", "install", "librtlsdr").Run()
}
