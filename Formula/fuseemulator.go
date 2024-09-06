package main

// FuseEmulator - Free Unix Spectrum Emulator
// Homepage: https://fuse-emulator.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installFuseEmulator() {
	// Método 1: Descargar y extraer .tar.gz
	fuseemulator_tar_url := "https://downloads.sourceforge.net/project/fuse-emulator/fuse/1.6.0/fuse-1.6.0.tar.gz"
	fuseemulator_cmd_tar := exec.Command("curl", "-L", fuseemulator_tar_url, "-o", "package.tar.gz")
	err := fuseemulator_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fuseemulator_zip_url := "https://downloads.sourceforge.net/project/fuse-emulator/fuse/1.6.0/fuse-1.6.0.zip"
	fuseemulator_cmd_zip := exec.Command("curl", "-L", fuseemulator_zip_url, "-o", "package.zip")
	err = fuseemulator_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fuseemulator_bin_url := "https://downloads.sourceforge.net/project/fuse-emulator/fuse/1.6.0/fuse-1.6.0.bin"
	fuseemulator_cmd_bin := exec.Command("curl", "-L", fuseemulator_bin_url, "-o", "binary.bin")
	err = fuseemulator_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fuseemulator_src_url := "https://downloads.sourceforge.net/project/fuse-emulator/fuse/1.6.0/fuse-1.6.0.src.tar.gz"
	fuseemulator_cmd_src := exec.Command("curl", "-L", fuseemulator_src_url, "-o", "source.tar.gz")
	err = fuseemulator_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fuseemulator_cmd_direct := exec.Command("./binary")
	err = fuseemulator_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libspectrum")
	exec.Command("latte", "install", "libspectrum").Run()
	fmt.Println("Instalando dependencia: sdl12-compat")
	exec.Command("latte", "install", "sdl12-compat").Run()
}
