package main

// Pwsafe - Generate passwords and manage encrypted password databases
// Homepage: https://github.com/nsd20463/pwsafe

import (
	"fmt"
	
	"os/exec"
)

func installPwsafe() {
	// Método 1: Descargar y extraer .tar.gz
	pwsafe_tar_url := "https://src.fedoraproject.org/repo/pkgs/pwsafe/pwsafe-0.2.0.tar.gz/4bb36538a2772ecbf1a542bc7d4746c0/pwsafe-0.2.0.tar.gz"
	pwsafe_cmd_tar := exec.Command("curl", "-L", pwsafe_tar_url, "-o", "package.tar.gz")
	err := pwsafe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pwsafe_zip_url := "https://src.fedoraproject.org/repo/pkgs/pwsafe/pwsafe-0.2.0.zip/4bb36538a2772ecbf1a542bc7d4746c0/pwsafe-0.2.0.zip"
	pwsafe_cmd_zip := exec.Command("curl", "-L", pwsafe_zip_url, "-o", "package.zip")
	err = pwsafe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pwsafe_bin_url := "https://src.fedoraproject.org/repo/pkgs/pwsafe/pwsafe-0.2.0.bin/4bb36538a2772ecbf1a542bc7d4746c0/pwsafe-0.2.0.bin"
	pwsafe_cmd_bin := exec.Command("curl", "-L", pwsafe_bin_url, "-o", "binary.bin")
	err = pwsafe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pwsafe_src_url := "https://src.fedoraproject.org/repo/pkgs/pwsafe/pwsafe-0.2.0.src.tar.gz/4bb36538a2772ecbf1a542bc7d4746c0/pwsafe-0.2.0.src.tar.gz"
	pwsafe_cmd_src := exec.Command("curl", "-L", pwsafe_src_url, "-o", "source.tar.gz")
	err = pwsafe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pwsafe_cmd_direct := exec.Command("./binary")
	err = pwsafe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
