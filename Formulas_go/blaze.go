package main

// Blaze - High-performance C++ math library for dense and sparse arithmetic
// Homepage: https://bitbucket.org/blaze-lib/blaze

import (
	"fmt"
	
	"os/exec"
)

func installBlaze() {
	// Método 1: Descargar y extraer .tar.gz
	blaze_tar_url := "https://bitbucket.org/blaze-lib/blaze/downloads/blaze-3.8.2.tar.gz"
	blaze_cmd_tar := exec.Command("curl", "-L", blaze_tar_url, "-o", "package.tar.gz")
	err := blaze_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	blaze_zip_url := "https://bitbucket.org/blaze-lib/blaze/downloads/blaze-3.8.2.zip"
	blaze_cmd_zip := exec.Command("curl", "-L", blaze_zip_url, "-o", "package.zip")
	err = blaze_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	blaze_bin_url := "https://bitbucket.org/blaze-lib/blaze/downloads/blaze-3.8.2.bin"
	blaze_cmd_bin := exec.Command("curl", "-L", blaze_bin_url, "-o", "binary.bin")
	err = blaze_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	blaze_src_url := "https://bitbucket.org/blaze-lib/blaze/downloads/blaze-3.8.2.src.tar.gz"
	blaze_cmd_src := exec.Command("curl", "-L", blaze_src_url, "-o", "source.tar.gz")
	err = blaze_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	blaze_cmd_direct := exec.Command("./binary")
	err = blaze_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: openblas")
exec.Command("latte", "install", "openblas")
}
