package main

// Supermodel - Sega Model 3 arcade emulator
// Homepage: https://www.supermodel3.com/

import (
	"fmt"
	
	"os/exec"
)

func installSupermodel() {
	// Método 1: Descargar y extraer .tar.gz
	supermodel_tar_url := "https://www.supermodel3.com/Files/Supermodel_0.2a_Src.zip"
	supermodel_cmd_tar := exec.Command("curl", "-L", supermodel_tar_url, "-o", "package.tar.gz")
	err := supermodel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	supermodel_zip_url := "https://www.supermodel3.com/Files/Supermodel_0.2a_Src.zip"
	supermodel_cmd_zip := exec.Command("curl", "-L", supermodel_zip_url, "-o", "package.zip")
	err = supermodel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	supermodel_bin_url := "https://www.supermodel3.com/Files/Supermodel_0.2a_Src.zip"
	supermodel_cmd_bin := exec.Command("curl", "-L", supermodel_bin_url, "-o", "binary.bin")
	err = supermodel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	supermodel_src_url := "https://www.supermodel3.com/Files/Supermodel_0.2a_Src.zip"
	supermodel_cmd_src := exec.Command("curl", "-L", supermodel_src_url, "-o", "source.tar.gz")
	err = supermodel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	supermodel_cmd_direct := exec.Command("./binary")
	err = supermodel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sdl12-compat")
exec.Command("latte", "install", "sdl12-compat")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
	fmt.Println("Instalando dependencia: mesa-glu")
exec.Command("latte", "install", "mesa-glu")
}
