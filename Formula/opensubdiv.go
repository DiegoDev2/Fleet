package main

// Opensubdiv - Open-source subdivision surface library
// Homepage: https://graphics.pixar.com/opensubdiv/docs/intro.html

import (
	"fmt"
	
	"os/exec"
)

func installOpensubdiv() {
	// Método 1: Descargar y extraer .tar.gz
	opensubdiv_tar_url := "https://github.com/PixarAnimationStudios/OpenSubdiv/archive/refs/tags/v3_6_0.tar.gz"
	opensubdiv_cmd_tar := exec.Command("curl", "-L", opensubdiv_tar_url, "-o", "package.tar.gz")
	err := opensubdiv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opensubdiv_zip_url := "https://github.com/PixarAnimationStudios/OpenSubdiv/archive/refs/tags/v3_6_0.zip"
	opensubdiv_cmd_zip := exec.Command("curl", "-L", opensubdiv_zip_url, "-o", "package.zip")
	err = opensubdiv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opensubdiv_bin_url := "https://github.com/PixarAnimationStudios/OpenSubdiv/archive/refs/tags/v3_6_0.bin"
	opensubdiv_cmd_bin := exec.Command("curl", "-L", opensubdiv_bin_url, "-o", "binary.bin")
	err = opensubdiv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opensubdiv_src_url := "https://github.com/PixarAnimationStudios/OpenSubdiv/archive/refs/tags/v3_6_0.src.tar.gz"
	opensubdiv_cmd_src := exec.Command("curl", "-L", opensubdiv_src_url, "-o", "source.tar.gz")
	err = opensubdiv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opensubdiv_cmd_direct := exec.Command("./binary")
	err = opensubdiv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: glfw")
	exec.Command("latte", "install", "glfw").Run()
}
