package main

// Nuget - Package manager for Microsoft development platform including .NET
// Homepage: https://www.nuget.org/

import (
	"fmt"
	
	"os/exec"
)

func installNuget() {
	// Método 1: Descargar y extraer .tar.gz
	nuget_tar_url := "https://dist.nuget.org/win-x86-commandline/v6.11.0/nuget.exe"
	nuget_cmd_tar := exec.Command("curl", "-L", nuget_tar_url, "-o", "package.tar.gz")
	err := nuget_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nuget_zip_url := "https://dist.nuget.org/win-x86-commandline/v6.11.0/nuget.exe"
	nuget_cmd_zip := exec.Command("curl", "-L", nuget_zip_url, "-o", "package.zip")
	err = nuget_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nuget_bin_url := "https://dist.nuget.org/win-x86-commandline/v6.11.0/nuget.exe"
	nuget_cmd_bin := exec.Command("curl", "-L", nuget_bin_url, "-o", "binary.bin")
	err = nuget_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nuget_src_url := "https://dist.nuget.org/win-x86-commandline/v6.11.0/nuget.exe"
	nuget_cmd_src := exec.Command("curl", "-L", nuget_src_url, "-o", "source.tar.gz")
	err = nuget_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nuget_cmd_direct := exec.Command("./binary")
	err = nuget_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: mono")
exec.Command("latte", "install", "mono")
}
