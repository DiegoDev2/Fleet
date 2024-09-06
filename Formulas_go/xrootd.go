package main

// Xrootd - High performance, scalable, fault-tolerant access to data
// Homepage: https://xrootd.slac.stanford.edu/

import (
	"fmt"
	
	"os/exec"
)

func installXrootd() {
	// Método 1: Descargar y extraer .tar.gz
	xrootd_tar_url := "https://github.com/xrootd/xrootd/releases/download/v5.7.0/xrootd-5.7.0.tar.gz"
	xrootd_cmd_tar := exec.Command("curl", "-L", xrootd_tar_url, "-o", "package.tar.gz")
	err := xrootd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xrootd_zip_url := "https://github.com/xrootd/xrootd/releases/download/v5.7.0/xrootd-5.7.0.zip"
	xrootd_cmd_zip := exec.Command("curl", "-L", xrootd_zip_url, "-o", "package.zip")
	err = xrootd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xrootd_bin_url := "https://github.com/xrootd/xrootd/releases/download/v5.7.0/xrootd-5.7.0.bin"
	xrootd_cmd_bin := exec.Command("curl", "-L", xrootd_bin_url, "-o", "binary.bin")
	err = xrootd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xrootd_src_url := "https://github.com/xrootd/xrootd/releases/download/v5.7.0/xrootd-5.7.0.src.tar.gz"
	xrootd_cmd_src := exec.Command("curl", "-L", xrootd_src_url, "-o", "source.tar.gz")
	err = xrootd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xrootd_cmd_direct := exec.Command("./binary")
	err = xrootd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: davix")
exec.Command("latte", "install", "davix")
	fmt.Println("Instalando dependencia: krb5")
exec.Command("latte", "install", "krb5")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
}
