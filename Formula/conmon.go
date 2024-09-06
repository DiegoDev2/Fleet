package main

// Conmon - OCI container runtime monitor
// Homepage: https://github.com/containers/conmon

import (
	"fmt"
	
	"os/exec"
)

func installConmon() {
	// Método 1: Descargar y extraer .tar.gz
	conmon_tar_url := "https://github.com/containers/conmon/archive/refs/tags/v2.1.12.tar.gz"
	conmon_cmd_tar := exec.Command("curl", "-L", conmon_tar_url, "-o", "package.tar.gz")
	err := conmon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	conmon_zip_url := "https://github.com/containers/conmon/archive/refs/tags/v2.1.12.zip"
	conmon_cmd_zip := exec.Command("curl", "-L", conmon_zip_url, "-o", "package.zip")
	err = conmon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	conmon_bin_url := "https://github.com/containers/conmon/archive/refs/tags/v2.1.12.bin"
	conmon_cmd_bin := exec.Command("curl", "-L", conmon_bin_url, "-o", "binary.bin")
	err = conmon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	conmon_src_url := "https://github.com/containers/conmon/archive/refs/tags/v2.1.12.src.tar.gz"
	conmon_cmd_src := exec.Command("curl", "-L", conmon_src_url, "-o", "source.tar.gz")
	err = conmon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	conmon_cmd_direct := exec.Command("./binary")
	err = conmon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: libseccomp")
	exec.Command("latte", "install", "libseccomp").Run()
	fmt.Println("Instalando dependencia: systemd")
	exec.Command("latte", "install", "systemd").Run()
}
