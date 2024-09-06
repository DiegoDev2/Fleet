package main

// Ftgl - Freetype / OpenGL bridge
// Homepage: https://sourceforge.net/projects/ftgl/

import (
	"fmt"
	
	"os/exec"
)

func installFtgl() {
	// Método 1: Descargar y extraer .tar.gz
	ftgl_tar_url := "https://downloads.sourceforge.net/project/ftgl/FTGL%20Source/2.1.3~rc5/ftgl-2.1.3-rc5.tar.gz"
	ftgl_cmd_tar := exec.Command("curl", "-L", ftgl_tar_url, "-o", "package.tar.gz")
	err := ftgl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ftgl_zip_url := "https://downloads.sourceforge.net/project/ftgl/FTGL%20Source/2.1.3~rc5/ftgl-2.1.3-rc5.zip"
	ftgl_cmd_zip := exec.Command("curl", "-L", ftgl_zip_url, "-o", "package.zip")
	err = ftgl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ftgl_bin_url := "https://downloads.sourceforge.net/project/ftgl/FTGL%20Source/2.1.3~rc5/ftgl-2.1.3-rc5.bin"
	ftgl_cmd_bin := exec.Command("curl", "-L", ftgl_bin_url, "-o", "binary.bin")
	err = ftgl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ftgl_src_url := "https://downloads.sourceforge.net/project/ftgl/FTGL%20Source/2.1.3~rc5/ftgl-2.1.3-rc5.src.tar.gz"
	ftgl_cmd_src := exec.Command("curl", "-L", ftgl_src_url, "-o", "source.tar.gz")
	err = ftgl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ftgl_cmd_direct := exec.Command("./binary")
	err = ftgl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
	fmt.Println("Instalando dependencia: mesa-glu")
exec.Command("latte", "install", "mesa-glu")
}
