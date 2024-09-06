package main

// Openttd - Simulation game based upon Transport Tycoon Deluxe
// Homepage: https://www.openttd.org/

import (
	"fmt"
	
	"os/exec"
)

func installOpenttd() {
	// Método 1: Descargar y extraer .tar.gz
	openttd_tar_url := "https://cdn.openttd.org/openttd-releases/14.1/openttd-14.1-source.tar.xz"
	openttd_cmd_tar := exec.Command("curl", "-L", openttd_tar_url, "-o", "package.tar.gz")
	err := openttd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openttd_zip_url := "https://cdn.openttd.org/openttd-releases/14.1/openttd-14.1-source.tar.xz"
	openttd_cmd_zip := exec.Command("curl", "-L", openttd_zip_url, "-o", "package.zip")
	err = openttd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openttd_bin_url := "https://cdn.openttd.org/openttd-releases/14.1/openttd-14.1-source.tar.xz"
	openttd_cmd_bin := exec.Command("curl", "-L", openttd_bin_url, "-o", "binary.bin")
	err = openttd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openttd_src_url := "https://cdn.openttd.org/openttd-releases/14.1/openttd-14.1-source.tar.xz"
	openttd_cmd_src := exec.Command("curl", "-L", openttd_src_url, "-o", "source.tar.gz")
	err = openttd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openttd_cmd_direct := exec.Command("./binary")
	err = openttd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: lzo")
	exec.Command("latte", "install", "lzo").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: fluid-synth")
	exec.Command("latte", "install", "fluid-synth").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: mesa-glu")
	exec.Command("latte", "install", "mesa-glu").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
}
