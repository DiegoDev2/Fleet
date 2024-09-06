package main

// LttngUst - Linux Trace Toolkit Next Generation Userspace Tracer
// Homepage: https://lttng.org/

import (
	"fmt"
	
	"os/exec"
)

func installLttngUst() {
	// Método 1: Descargar y extraer .tar.gz
	lttngust_tar_url := "https://lttng.org/files/lttng-ust/lttng-ust-2.13.8.tar.bz2"
	lttngust_cmd_tar := exec.Command("curl", "-L", lttngust_tar_url, "-o", "package.tar.gz")
	err := lttngust_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lttngust_zip_url := "https://lttng.org/files/lttng-ust/lttng-ust-2.13.8.tar.bz2"
	lttngust_cmd_zip := exec.Command("curl", "-L", lttngust_zip_url, "-o", "package.zip")
	err = lttngust_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lttngust_bin_url := "https://lttng.org/files/lttng-ust/lttng-ust-2.13.8.tar.bz2"
	lttngust_cmd_bin := exec.Command("curl", "-L", lttngust_bin_url, "-o", "binary.bin")
	err = lttngust_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lttngust_src_url := "https://lttng.org/files/lttng-ust/lttng-ust-2.13.8.tar.bz2"
	lttngust_cmd_src := exec.Command("curl", "-L", lttngust_src_url, "-o", "source.tar.gz")
	err = lttngust_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lttngust_cmd_direct := exec.Command("./binary")
	err = lttngust_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: numactl")
exec.Command("latte", "install", "numactl")
	fmt.Println("Instalando dependencia: userspace-rcu")
exec.Command("latte", "install", "userspace-rcu")
}
