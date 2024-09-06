package main

// Hwloc - Portable abstraction of the hierarchical topology of modern architectures
// Homepage: https://www.open-mpi.org/projects/hwloc/

import (
	"fmt"
	
	"os/exec"
)

func installHwloc() {
	// Método 1: Descargar y extraer .tar.gz
	hwloc_tar_url := "https://download.open-mpi.org/release/hwloc/v2.11/hwloc-2.11.1.tar.bz2"
	hwloc_cmd_tar := exec.Command("curl", "-L", hwloc_tar_url, "-o", "package.tar.gz")
	err := hwloc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hwloc_zip_url := "https://download.open-mpi.org/release/hwloc/v2.11/hwloc-2.11.1.tar.bz2"
	hwloc_cmd_zip := exec.Command("curl", "-L", hwloc_zip_url, "-o", "package.zip")
	err = hwloc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hwloc_bin_url := "https://download.open-mpi.org/release/hwloc/v2.11/hwloc-2.11.1.tar.bz2"
	hwloc_cmd_bin := exec.Command("curl", "-L", hwloc_bin_url, "-o", "binary.bin")
	err = hwloc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hwloc_src_url := "https://download.open-mpi.org/release/hwloc/v2.11/hwloc-2.11.1.tar.bz2"
	hwloc_cmd_src := exec.Command("curl", "-L", hwloc_src_url, "-o", "source.tar.gz")
	err = hwloc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hwloc_cmd_direct := exec.Command("./binary")
	err = hwloc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
