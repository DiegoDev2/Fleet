package main

// DotnetAT6 - .NET Core
// Homepage: https://dotnet.microsoft.com/

import (
	"fmt"
	
	"os/exec"
)

func installDotnetAT6() {
	// Método 1: Descargar y extraer .tar.gz
	dotnetat6_tar_url := "https://github.com/dotnet/installer.git"
	dotnetat6_cmd_tar := exec.Command("curl", "-L", dotnetat6_tar_url, "-o", "package.tar.gz")
	err := dotnetat6_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dotnetat6_zip_url := "https://github.com/dotnet/installer.git"
	dotnetat6_cmd_zip := exec.Command("curl", "-L", dotnetat6_zip_url, "-o", "package.zip")
	err = dotnetat6_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dotnetat6_bin_url := "https://github.com/dotnet/installer.git"
	dotnetat6_cmd_bin := exec.Command("curl", "-L", dotnetat6_bin_url, "-o", "binary.bin")
	err = dotnetat6_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dotnetat6_src_url := "https://github.com/dotnet/installer.git"
	dotnetat6_cmd_src := exec.Command("curl", "-L", dotnetat6_src_url, "-o", "source.tar.gz")
	err = dotnetat6_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dotnetat6_cmd_direct := exec.Command("./binary")
	err = dotnetat6_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: libunwind")
exec.Command("latte", "install", "libunwind")
	fmt.Println("Instalando dependencia: lttng-ust")
exec.Command("latte", "install", "lttng-ust")
}
