package main

// Gphoto2 - Command-line interface to libgphoto2
// Homepage: http://www.gphoto.org/

import (
	"fmt"
	
	"os/exec"
)

func installGphoto2() {
	// Método 1: Descargar y extraer .tar.gz
	gphoto2_tar_url := "https://downloads.sourceforge.net/project/gphoto/gphoto/2.5.28/gphoto2-2.5.28.tar.bz2"
	gphoto2_cmd_tar := exec.Command("curl", "-L", gphoto2_tar_url, "-o", "package.tar.gz")
	err := gphoto2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gphoto2_zip_url := "https://downloads.sourceforge.net/project/gphoto/gphoto/2.5.28/gphoto2-2.5.28.tar.bz2"
	gphoto2_cmd_zip := exec.Command("curl", "-L", gphoto2_zip_url, "-o", "package.zip")
	err = gphoto2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gphoto2_bin_url := "https://downloads.sourceforge.net/project/gphoto/gphoto/2.5.28/gphoto2-2.5.28.tar.bz2"
	gphoto2_cmd_bin := exec.Command("curl", "-L", gphoto2_bin_url, "-o", "binary.bin")
	err = gphoto2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gphoto2_src_url := "https://downloads.sourceforge.net/project/gphoto/gphoto/2.5.28/gphoto2-2.5.28.tar.bz2"
	gphoto2_cmd_src := exec.Command("curl", "-L", gphoto2_src_url, "-o", "source.tar.gz")
	err = gphoto2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gphoto2_cmd_direct := exec.Command("./binary")
	err = gphoto2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libexif")
	exec.Command("latte", "install", "libexif").Run()
	fmt.Println("Instalando dependencia: libgphoto2")
	exec.Command("latte", "install", "libgphoto2").Run()
	fmt.Println("Instalando dependencia: popt")
	exec.Command("latte", "install", "popt").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
