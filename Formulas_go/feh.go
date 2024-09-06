package main

// Feh - X11 image viewer
// Homepage: https://feh.finalrewind.org/

import (
	"fmt"
	
	"os/exec"
)

func installFeh() {
	// Método 1: Descargar y extraer .tar.gz
	feh_tar_url := "https://feh.finalrewind.org/feh-3.10.3.tar.bz2"
	feh_cmd_tar := exec.Command("curl", "-L", feh_tar_url, "-o", "package.tar.gz")
	err := feh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	feh_zip_url := "https://feh.finalrewind.org/feh-3.10.3.tar.bz2"
	feh_cmd_zip := exec.Command("curl", "-L", feh_zip_url, "-o", "package.zip")
	err = feh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	feh_bin_url := "https://feh.finalrewind.org/feh-3.10.3.tar.bz2"
	feh_cmd_bin := exec.Command("curl", "-L", feh_bin_url, "-o", "binary.bin")
	err = feh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	feh_src_url := "https://feh.finalrewind.org/feh-3.10.3.tar.bz2"
	feh_cmd_src := exec.Command("curl", "-L", feh_src_url, "-o", "source.tar.gz")
	err = feh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	feh_cmd_direct := exec.Command("./binary")
	err = feh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: imlib2")
exec.Command("latte", "install", "imlib2")
	fmt.Println("Instalando dependencia: libexif")
exec.Command("latte", "install", "libexif")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxinerama")
exec.Command("latte", "install", "libxinerama")
	fmt.Println("Instalando dependencia: libxt")
exec.Command("latte", "install", "libxt")
}
