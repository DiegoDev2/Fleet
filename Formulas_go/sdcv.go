package main

// Sdcv - StarDict Console Version
// Homepage: https://dushistov.github.io/sdcv/

import (
	"fmt"
	
	"os/exec"
)

func installSdcv() {
	// Método 1: Descargar y extraer .tar.gz
	sdcv_tar_url := "https://github.com/Dushistov/sdcv/archive/refs/tags/v0.5.5.tar.gz"
	sdcv_cmd_tar := exec.Command("curl", "-L", sdcv_tar_url, "-o", "package.tar.gz")
	err := sdcv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sdcv_zip_url := "https://github.com/Dushistov/sdcv/archive/refs/tags/v0.5.5.zip"
	sdcv_cmd_zip := exec.Command("curl", "-L", sdcv_zip_url, "-o", "package.zip")
	err = sdcv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sdcv_bin_url := "https://github.com/Dushistov/sdcv/archive/refs/tags/v0.5.5.bin"
	sdcv_cmd_bin := exec.Command("curl", "-L", sdcv_bin_url, "-o", "binary.bin")
	err = sdcv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sdcv_src_url := "https://github.com/Dushistov/sdcv/archive/refs/tags/v0.5.5.src.tar.gz"
	sdcv_cmd_src := exec.Command("curl", "-L", sdcv_src_url, "-o", "source.tar.gz")
	err = sdcv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sdcv_cmd_direct := exec.Command("./binary")
	err = sdcv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
