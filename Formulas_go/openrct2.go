package main

// Openrct2 - Open source re-implementation of RollerCoaster Tycoon 2
// Homepage: https://openrct2.io/

import (
	"fmt"
	
	"os/exec"
)

func installOpenrct2() {
	// Método 1: Descargar y extraer .tar.gz
	openrct2_tar_url := "https://github.com/OpenRCT2/OpenRCT2.git"
	openrct2_cmd_tar := exec.Command("curl", "-L", openrct2_tar_url, "-o", "package.tar.gz")
	err := openrct2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openrct2_zip_url := "https://github.com/OpenRCT2/OpenRCT2.git"
	openrct2_cmd_zip := exec.Command("curl", "-L", openrct2_zip_url, "-o", "package.zip")
	err = openrct2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openrct2_bin_url := "https://github.com/OpenRCT2/OpenRCT2.git"
	openrct2_cmd_bin := exec.Command("curl", "-L", openrct2_bin_url, "-o", "binary.bin")
	err = openrct2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openrct2_src_url := "https://github.com/OpenRCT2/OpenRCT2.git"
	openrct2_cmd_src := exec.Command("curl", "-L", openrct2_src_url, "-o", "source.tar.gz")
	err = openrct2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openrct2_cmd_direct := exec.Command("./binary")
	err = openrct2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: nlohmann-json")
exec.Command("latte", "install", "nlohmann-json")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: duktape")
exec.Command("latte", "install", "duktape")
	fmt.Println("Instalando dependencia: flac")
exec.Command("latte", "install", "flac")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: libogg")
exec.Command("latte", "install", "libogg")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libvorbis")
exec.Command("latte", "install", "libvorbis")
	fmt.Println("Instalando dependencia: libzip")
exec.Command("latte", "install", "libzip")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: speexdsp")
exec.Command("latte", "install", "speexdsp")
	fmt.Println("Instalando dependencia: curl")
exec.Command("latte", "install", "curl")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
}
