package main

// Bochs - Open source IA-32 (x86) PC emulator written in C++
// Homepage: https://bochs.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installBochs() {
	// Método 1: Descargar y extraer .tar.gz
	bochs_tar_url := "https://downloads.sourceforge.net/project/bochs/bochs/2.8/bochs-2.8.tar.gz"
	bochs_cmd_tar := exec.Command("curl", "-L", bochs_tar_url, "-o", "package.tar.gz")
	err := bochs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bochs_zip_url := "https://downloads.sourceforge.net/project/bochs/bochs/2.8/bochs-2.8.zip"
	bochs_cmd_zip := exec.Command("curl", "-L", bochs_zip_url, "-o", "package.zip")
	err = bochs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bochs_bin_url := "https://downloads.sourceforge.net/project/bochs/bochs/2.8/bochs-2.8.bin"
	bochs_cmd_bin := exec.Command("curl", "-L", bochs_bin_url, "-o", "binary.bin")
	err = bochs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bochs_src_url := "https://downloads.sourceforge.net/project/bochs/bochs/2.8/bochs-2.8.src.tar.gz"
	bochs_cmd_src := exec.Command("curl", "-L", bochs_src_url, "-o", "source.tar.gz")
	err = bochs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bochs_cmd_direct := exec.Command("./binary")
	err = bochs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
