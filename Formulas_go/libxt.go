package main

// Libxt - X.Org: X Toolkit Intrinsics library
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibxt() {
	// Método 1: Descargar y extraer .tar.gz
	libxt_tar_url := "https://www.x.org/archive/individual/lib/libXt-1.3.0.tar.xz"
	libxt_cmd_tar := exec.Command("curl", "-L", libxt_tar_url, "-o", "package.tar.gz")
	err := libxt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxt_zip_url := "https://www.x.org/archive/individual/lib/libXt-1.3.0.tar.xz"
	libxt_cmd_zip := exec.Command("curl", "-L", libxt_zip_url, "-o", "package.zip")
	err = libxt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxt_bin_url := "https://www.x.org/archive/individual/lib/libXt-1.3.0.tar.xz"
	libxt_cmd_bin := exec.Command("curl", "-L", libxt_bin_url, "-o", "binary.bin")
	err = libxt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxt_src_url := "https://www.x.org/archive/individual/lib/libXt-1.3.0.tar.xz"
	libxt_cmd_src := exec.Command("curl", "-L", libxt_src_url, "-o", "source.tar.gz")
	err = libxt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxt_cmd_direct := exec.Command("./binary")
	err = libxt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libice")
exec.Command("latte", "install", "libice")
	fmt.Println("Instalando dependencia: libsm")
exec.Command("latte", "install", "libsm")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
}
