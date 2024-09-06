package main

// ScummvmTools - Collection of tools for ScummVM
// Homepage: https://www.scummvm.org/

import (
	"fmt"
	
	"os/exec"
)

func installScummvmTools() {
	// Método 1: Descargar y extraer .tar.gz
	scummvmtools_tar_url := "https://downloads.scummvm.org/frs/scummvm-tools/2.7.0/scummvm-tools-2.7.0.tar.xz"
	scummvmtools_cmd_tar := exec.Command("curl", "-L", scummvmtools_tar_url, "-o", "package.tar.gz")
	err := scummvmtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scummvmtools_zip_url := "https://downloads.scummvm.org/frs/scummvm-tools/2.7.0/scummvm-tools-2.7.0.tar.xz"
	scummvmtools_cmd_zip := exec.Command("curl", "-L", scummvmtools_zip_url, "-o", "package.zip")
	err = scummvmtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scummvmtools_bin_url := "https://downloads.scummvm.org/frs/scummvm-tools/2.7.0/scummvm-tools-2.7.0.tar.xz"
	scummvmtools_cmd_bin := exec.Command("curl", "-L", scummvmtools_bin_url, "-o", "binary.bin")
	err = scummvmtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scummvmtools_src_url := "https://downloads.scummvm.org/frs/scummvm-tools/2.7.0/scummvm-tools-2.7.0.tar.xz"
	scummvmtools_cmd_src := exec.Command("curl", "-L", scummvmtools_src_url, "-o", "source.tar.gz")
	err = scummvmtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scummvmtools_cmd_direct := exec.Command("./binary")
	err = scummvmtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: flac")
	exec.Command("latte", "install", "flac").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: libogg")
	exec.Command("latte", "install", "libogg").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
	fmt.Println("Instalando dependencia: mad")
	exec.Command("latte", "install", "mad").Run()
	fmt.Println("Instalando dependencia: wxwidgets")
	exec.Command("latte", "install", "wxwidgets").Run()
}
