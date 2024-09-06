package main

// C10t - Minecraft cartography tool
// Homepage: https://github.com/udoprog/c10t

import (
	"fmt"
	
	"os/exec"
)

func installC10t() {
	// Método 1: Descargar y extraer .tar.gz
	c10t_tar_url := "https://github.com/udoprog/c10t/archive/refs/tags/1.7.tar.gz"
	c10t_cmd_tar := exec.Command("curl", "-L", c10t_tar_url, "-o", "package.tar.gz")
	err := c10t_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	c10t_zip_url := "https://github.com/udoprog/c10t/archive/refs/tags/1.7.zip"
	c10t_cmd_zip := exec.Command("curl", "-L", c10t_zip_url, "-o", "package.zip")
	err = c10t_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	c10t_bin_url := "https://github.com/udoprog/c10t/archive/refs/tags/1.7.bin"
	c10t_cmd_bin := exec.Command("curl", "-L", c10t_bin_url, "-o", "binary.bin")
	err = c10t_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	c10t_src_url := "https://github.com/udoprog/c10t/archive/refs/tags/1.7.src.tar.gz"
	c10t_cmd_src := exec.Command("curl", "-L", c10t_src_url, "-o", "source.tar.gz")
	err = c10t_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	c10t_cmd_direct := exec.Command("./binary")
	err = c10t_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
}
