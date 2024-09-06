package main

// Gl2ps - OpenGL to PostScript printing library
// Homepage: https://www.geuz.org/gl2ps/

import (
	"fmt"
	
	"os/exec"
)

func installGl2ps() {
	// Método 1: Descargar y extraer .tar.gz
	gl2ps_tar_url := "https://geuz.org/gl2ps/src/gl2ps-1.4.2.tgz"
	gl2ps_cmd_tar := exec.Command("curl", "-L", gl2ps_tar_url, "-o", "package.tar.gz")
	err := gl2ps_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gl2ps_zip_url := "https://geuz.org/gl2ps/src/gl2ps-1.4.2.tgz"
	gl2ps_cmd_zip := exec.Command("curl", "-L", gl2ps_zip_url, "-o", "package.zip")
	err = gl2ps_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gl2ps_bin_url := "https://geuz.org/gl2ps/src/gl2ps-1.4.2.tgz"
	gl2ps_cmd_bin := exec.Command("curl", "-L", gl2ps_bin_url, "-o", "binary.bin")
	err = gl2ps_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gl2ps_src_url := "https://geuz.org/gl2ps/src/gl2ps-1.4.2.tgz"
	gl2ps_cmd_src := exec.Command("curl", "-L", gl2ps_src_url, "-o", "source.tar.gz")
	err = gl2ps_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gl2ps_cmd_direct := exec.Command("./binary")
	err = gl2ps_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: freeglut")
	exec.Command("latte", "install", "freeglut").Run()
}
