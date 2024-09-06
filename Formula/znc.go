package main

// Znc - Advanced IRC bouncer
// Homepage: https://wiki.znc.in/ZNC

import (
	"fmt"
	
	"os/exec"
)

func installZnc() {
	// Método 1: Descargar y extraer .tar.gz
	znc_tar_url := "https://znc.in/releases/znc-1.9.1.tar.gz"
	znc_cmd_tar := exec.Command("curl", "-L", znc_tar_url, "-o", "package.tar.gz")
	err := znc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	znc_zip_url := "https://znc.in/releases/znc-1.9.1.zip"
	znc_cmd_zip := exec.Command("curl", "-L", znc_zip_url, "-o", "package.zip")
	err = znc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	znc_bin_url := "https://znc.in/releases/znc-1.9.1.bin"
	znc_cmd_bin := exec.Command("curl", "-L", znc_bin_url, "-o", "binary.bin")
	err = znc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	znc_src_url := "https://znc.in/releases/znc-1.9.1.src.tar.gz"
	znc_cmd_src := exec.Command("curl", "-L", znc_src_url, "-o", "source.tar.gz")
	err = znc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	znc_cmd_direct := exec.Command("./binary")
	err = znc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
