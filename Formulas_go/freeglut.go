package main

// Freeglut - Open-source alternative to the OpenGL Utility Toolkit (GLUT) library
// Homepage: https://freeglut.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installFreeglut() {
	// Método 1: Descargar y extraer .tar.gz
	freeglut_tar_url := "https://github.com/FreeGLUTProject/freeglut/releases/download/v3.6.0/freeglut-3.6.0.tar.gz"
	freeglut_cmd_tar := exec.Command("curl", "-L", freeglut_tar_url, "-o", "package.tar.gz")
	err := freeglut_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	freeglut_zip_url := "https://github.com/FreeGLUTProject/freeglut/releases/download/v3.6.0/freeglut-3.6.0.zip"
	freeglut_cmd_zip := exec.Command("curl", "-L", freeglut_zip_url, "-o", "package.zip")
	err = freeglut_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	freeglut_bin_url := "https://github.com/FreeGLUTProject/freeglut/releases/download/v3.6.0/freeglut-3.6.0.bin"
	freeglut_cmd_bin := exec.Command("curl", "-L", freeglut_bin_url, "-o", "binary.bin")
	err = freeglut_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	freeglut_src_url := "https://github.com/FreeGLUTProject/freeglut/releases/download/v3.6.0/freeglut-3.6.0.src.tar.gz"
	freeglut_cmd_src := exec.Command("curl", "-L", freeglut_src_url, "-o", "source.tar.gz")
	err = freeglut_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	freeglut_cmd_direct := exec.Command("./binary")
	err = freeglut_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxi")
exec.Command("latte", "install", "libxi")
	fmt.Println("Instalando dependencia: libxrandr")
exec.Command("latte", "install", "libxrandr")
	fmt.Println("Instalando dependencia: libxxf86vm")
exec.Command("latte", "install", "libxxf86vm")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
	fmt.Println("Instalando dependencia: mesa-glu")
exec.Command("latte", "install", "mesa-glu")
	fmt.Println("Instalando dependencia: xinput")
exec.Command("latte", "install", "xinput")
}
