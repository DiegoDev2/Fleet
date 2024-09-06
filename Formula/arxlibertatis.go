package main

// ArxLibertatis - Cross-platform, open source port of Arx Fatalis
// Homepage: https://arx-libertatis.org/

import (
	"fmt"
	
	"os/exec"
)

func installArxLibertatis() {
	// Método 1: Descargar y extraer .tar.gz
	arxlibertatis_tar_url := "https://arx-libertatis.org/files/arx-libertatis-1.2.1/arx-libertatis-1.2.1.tar.xz"
	arxlibertatis_cmd_tar := exec.Command("curl", "-L", arxlibertatis_tar_url, "-o", "package.tar.gz")
	err := arxlibertatis_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	arxlibertatis_zip_url := "https://arx-libertatis.org/files/arx-libertatis-1.2.1/arx-libertatis-1.2.1.tar.xz"
	arxlibertatis_cmd_zip := exec.Command("curl", "-L", arxlibertatis_zip_url, "-o", "package.zip")
	err = arxlibertatis_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	arxlibertatis_bin_url := "https://arx-libertatis.org/files/arx-libertatis-1.2.1/arx-libertatis-1.2.1.tar.xz"
	arxlibertatis_cmd_bin := exec.Command("curl", "-L", arxlibertatis_bin_url, "-o", "binary.bin")
	err = arxlibertatis_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	arxlibertatis_src_url := "https://arx-libertatis.org/files/arx-libertatis-1.2.1/arx-libertatis-1.2.1.tar.xz"
	arxlibertatis_cmd_src := exec.Command("curl", "-L", arxlibertatis_src_url, "-o", "source.tar.gz")
	err = arxlibertatis_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	arxlibertatis_cmd_direct := exec.Command("./binary")
	err = arxlibertatis_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: glm")
	exec.Command("latte", "install", "glm").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: glew")
	exec.Command("latte", "install", "glew").Run()
	fmt.Println("Instalando dependencia: innoextract")
	exec.Command("latte", "install", "innoextract").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: openal-soft")
	exec.Command("latte", "install", "openal-soft").Run()
}
