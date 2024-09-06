package main

// Simutrans - Transport simulator
// Homepage: https://www.simutrans.com/

import (
	"fmt"
	
	"os/exec"
)

func installSimutrans() {
	// Método 1: Descargar y extraer .tar.gz
	simutrans_tar_url := "svn://servers.simutrans.org/simutrans/trunk/"
	simutrans_cmd_tar := exec.Command("curl", "-L", simutrans_tar_url, "-o", "package.tar.gz")
	err := simutrans_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	simutrans_zip_url := "svn://servers.simutrans.org/simutrans/trunk/"
	simutrans_cmd_zip := exec.Command("curl", "-L", simutrans_zip_url, "-o", "package.zip")
	err = simutrans_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	simutrans_bin_url := "svn://servers.simutrans.org/simutrans/trunk/"
	simutrans_cmd_bin := exec.Command("curl", "-L", simutrans_bin_url, "-o", "binary.bin")
	err = simutrans_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	simutrans_src_url := "svn://servers.simutrans.org/simutrans/trunk/"
	simutrans_cmd_src := exec.Command("curl", "-L", simutrans_src_url, "-o", "source.tar.gz")
	err = simutrans_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	simutrans_cmd_direct := exec.Command("./binary")
	err = simutrans_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: fluid-synth")
exec.Command("latte", "install", "fluid-synth")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: miniupnpc")
exec.Command("latte", "install", "miniupnpc")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
