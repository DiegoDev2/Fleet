package main

// FabricInstaller - Installer for Fabric for the vanilla launcher
// Homepage: https://fabricmc.net/

import (
	"fmt"
	
	"os/exec"
)

func installFabricInstaller() {
	// Método 1: Descargar y extraer .tar.gz
	fabricinstaller_tar_url := "https://maven.fabricmc.net/net/fabricmc/fabric-installer/1.0.1/fabric-installer-1.0.1.jar"
	fabricinstaller_cmd_tar := exec.Command("curl", "-L", fabricinstaller_tar_url, "-o", "package.tar.gz")
	err := fabricinstaller_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fabricinstaller_zip_url := "https://maven.fabricmc.net/net/fabricmc/fabric-installer/1.0.1/fabric-installer-1.0.1.jar"
	fabricinstaller_cmd_zip := exec.Command("curl", "-L", fabricinstaller_zip_url, "-o", "package.zip")
	err = fabricinstaller_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fabricinstaller_bin_url := "https://maven.fabricmc.net/net/fabricmc/fabric-installer/1.0.1/fabric-installer-1.0.1.jar"
	fabricinstaller_cmd_bin := exec.Command("curl", "-L", fabricinstaller_bin_url, "-o", "binary.bin")
	err = fabricinstaller_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fabricinstaller_src_url := "https://maven.fabricmc.net/net/fabricmc/fabric-installer/1.0.1/fabric-installer-1.0.1.jar"
	fabricinstaller_cmd_src := exec.Command("curl", "-L", fabricinstaller_src_url, "-o", "source.tar.gz")
	err = fabricinstaller_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fabricinstaller_cmd_direct := exec.Command("./binary")
	err = fabricinstaller_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
