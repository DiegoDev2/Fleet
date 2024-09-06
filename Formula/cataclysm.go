package main

// Cataclysm - Fork/variant of Cataclysm Roguelike
// Homepage: https://github.com/CleverRaven/Cataclysm-DDA

import (
	"fmt"
	
	"os/exec"
)

func installCataclysm() {
	// Método 1: Descargar y extraer .tar.gz
	cataclysm_tar_url := "https://github.com/CleverRaven/Cataclysm-DDA/archive/refs/tags/0.G.tar.gz"
	cataclysm_cmd_tar := exec.Command("curl", "-L", cataclysm_tar_url, "-o", "package.tar.gz")
	err := cataclysm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cataclysm_zip_url := "https://github.com/CleverRaven/Cataclysm-DDA/archive/refs/tags/0.G.zip"
	cataclysm_cmd_zip := exec.Command("curl", "-L", cataclysm_zip_url, "-o", "package.zip")
	err = cataclysm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cataclysm_bin_url := "https://github.com/CleverRaven/Cataclysm-DDA/archive/refs/tags/0.G.bin"
	cataclysm_cmd_bin := exec.Command("curl", "-L", cataclysm_bin_url, "-o", "binary.bin")
	err = cataclysm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cataclysm_src_url := "https://github.com/CleverRaven/Cataclysm-DDA/archive/refs/tags/0.G.src.tar.gz"
	cataclysm_cmd_src := exec.Command("curl", "-L", cataclysm_src_url, "-o", "source.tar.gz")
	err = cataclysm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cataclysm_cmd_direct := exec.Command("./binary")
	err = cataclysm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libogg")
	exec.Command("latte", "install", "libogg").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: sdl2_image")
	exec.Command("latte", "install", "sdl2_image").Run()
	fmt.Println("Instalando dependencia: sdl2_mixer")
	exec.Command("latte", "install", "sdl2_mixer").Run()
	fmt.Println("Instalando dependencia: sdl2_ttf")
	exec.Command("latte", "install", "sdl2_ttf").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
