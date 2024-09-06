package main

// Libxvmc - X.Org: X-Video Motion Compensation API
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibxvmc() {
	// Método 1: Descargar y extraer .tar.gz
	libxvmc_tar_url := "https://www.x.org/archive/individual/lib/libXvMC-1.0.14.tar.xz"
	libxvmc_cmd_tar := exec.Command("curl", "-L", libxvmc_tar_url, "-o", "package.tar.gz")
	err := libxvmc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxvmc_zip_url := "https://www.x.org/archive/individual/lib/libXvMC-1.0.14.tar.xz"
	libxvmc_cmd_zip := exec.Command("curl", "-L", libxvmc_zip_url, "-o", "package.zip")
	err = libxvmc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxvmc_bin_url := "https://www.x.org/archive/individual/lib/libXvMC-1.0.14.tar.xz"
	libxvmc_cmd_bin := exec.Command("curl", "-L", libxvmc_bin_url, "-o", "binary.bin")
	err = libxvmc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxvmc_src_url := "https://www.x.org/archive/individual/lib/libXvMC-1.0.14.tar.xz"
	libxvmc_cmd_src := exec.Command("curl", "-L", libxvmc_src_url, "-o", "source.tar.gz")
	err = libxvmc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxvmc_cmd_direct := exec.Command("./binary")
	err = libxvmc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: util-macros")
exec.Command("latte", "install", "util-macros")
	fmt.Println("Instalando dependencia: xorgproto")
exec.Command("latte", "install", "xorgproto")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxext")
exec.Command("latte", "install", "libxext")
	fmt.Println("Instalando dependencia: libxv")
exec.Command("latte", "install", "libxv")
}
