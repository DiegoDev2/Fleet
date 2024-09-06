package main

// Sfml - Multi-media library with bindings for multiple languages
// Homepage: https://www.sfml-dev.org/

import (
	"fmt"
	
	"os/exec"
)

func installSfml() {
	// Método 1: Descargar y extraer .tar.gz
	sfml_tar_url := "https://www.sfml-dev.org/files/SFML-2.6.1-sources.zip"
	sfml_cmd_tar := exec.Command("curl", "-L", sfml_tar_url, "-o", "package.tar.gz")
	err := sfml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sfml_zip_url := "https://www.sfml-dev.org/files/SFML-2.6.1-sources.zip"
	sfml_cmd_zip := exec.Command("curl", "-L", sfml_zip_url, "-o", "package.zip")
	err = sfml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sfml_bin_url := "https://www.sfml-dev.org/files/SFML-2.6.1-sources.zip"
	sfml_cmd_bin := exec.Command("curl", "-L", sfml_bin_url, "-o", "binary.bin")
	err = sfml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sfml_src_url := "https://www.sfml-dev.org/files/SFML-2.6.1-sources.zip"
	sfml_cmd_src := exec.Command("curl", "-L", sfml_src_url, "-o", "source.tar.gz")
	err = sfml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sfml_cmd_direct := exec.Command("./binary")
	err = sfml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: doxygen")
	exec.Command("latte", "install", "doxygen").Run()
	fmt.Println("Instalando dependencia: flac")
	exec.Command("latte", "install", "flac").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: libogg")
	exec.Command("latte", "install", "libogg").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxcursor")
	exec.Command("latte", "install", "libxcursor").Run()
	fmt.Println("Instalando dependencia: libxrandr")
	exec.Command("latte", "install", "libxrandr").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: mesa-glu")
	exec.Command("latte", "install", "mesa-glu").Run()
	fmt.Println("Instalando dependencia: openal-soft")
	exec.Command("latte", "install", "openal-soft").Run()
	fmt.Println("Instalando dependencia: systemd")
	exec.Command("latte", "install", "systemd").Run()
}
