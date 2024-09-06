package main

// Ode - Simulating articulated rigid body dynamics
// Homepage: https://www.ode.org/

import (
	"fmt"
	
	"os/exec"
)

func installOde() {
	// Método 1: Descargar y extraer .tar.gz
	ode_tar_url := "https://bitbucket.org/odedevs/ode/downloads/ode-0.16.5.tar.gz"
	ode_cmd_tar := exec.Command("curl", "-L", ode_tar_url, "-o", "package.tar.gz")
	err := ode_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ode_zip_url := "https://bitbucket.org/odedevs/ode/downloads/ode-0.16.5.zip"
	ode_cmd_zip := exec.Command("curl", "-L", ode_zip_url, "-o", "package.zip")
	err = ode_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ode_bin_url := "https://bitbucket.org/odedevs/ode/downloads/ode-0.16.5.bin"
	ode_cmd_bin := exec.Command("curl", "-L", ode_bin_url, "-o", "binary.bin")
	err = ode_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ode_src_url := "https://bitbucket.org/odedevs/ode/downloads/ode-0.16.5.src.tar.gz"
	ode_cmd_src := exec.Command("curl", "-L", ode_src_url, "-o", "source.tar.gz")
	err = ode_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ode_cmd_direct := exec.Command("./binary")
	err = ode_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libccd")
exec.Command("latte", "install", "libccd")
}
