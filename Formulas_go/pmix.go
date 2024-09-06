package main

// Pmix - Process Management Interface for HPC environments
// Homepage: https://openpmix.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installPmix() {
	// Método 1: Descargar y extraer .tar.gz
	pmix_tar_url := "https://github.com/openpmix/openpmix/releases/download/v5.0.3/pmix-5.0.3.tar.bz2"
	pmix_cmd_tar := exec.Command("curl", "-L", pmix_tar_url, "-o", "package.tar.gz")
	err := pmix_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pmix_zip_url := "https://github.com/openpmix/openpmix/releases/download/v5.0.3/pmix-5.0.3.tar.bz2"
	pmix_cmd_zip := exec.Command("curl", "-L", pmix_zip_url, "-o", "package.zip")
	err = pmix_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pmix_bin_url := "https://github.com/openpmix/openpmix/releases/download/v5.0.3/pmix-5.0.3.tar.bz2"
	pmix_cmd_bin := exec.Command("curl", "-L", pmix_bin_url, "-o", "binary.bin")
	err = pmix_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pmix_src_url := "https://github.com/openpmix/openpmix/releases/download/v5.0.3/pmix-5.0.3.tar.bz2"
	pmix_cmd_src := exec.Command("curl", "-L", pmix_src_url, "-o", "source.tar.gz")
	err = pmix_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pmix_cmd_direct := exec.Command("./binary")
	err = pmix_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: hwloc")
exec.Command("latte", "install", "hwloc")
	fmt.Println("Instalando dependencia: libevent")
exec.Command("latte", "install", "libevent")
}
