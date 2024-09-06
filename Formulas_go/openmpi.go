package main

// OpenMpi - High performance message passing library
// Homepage: https://www.open-mpi.org/

import (
	"fmt"
	
	"os/exec"
)

func installOpenMpi() {
	// Método 1: Descargar y extraer .tar.gz
	openmpi_tar_url := "https://download.open-mpi.org/release/open-mpi/v5.0/openmpi-5.0.3.tar.bz2"
	openmpi_cmd_tar := exec.Command("curl", "-L", openmpi_tar_url, "-o", "package.tar.gz")
	err := openmpi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openmpi_zip_url := "https://download.open-mpi.org/release/open-mpi/v5.0/openmpi-5.0.3.tar.bz2"
	openmpi_cmd_zip := exec.Command("curl", "-L", openmpi_zip_url, "-o", "package.zip")
	err = openmpi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openmpi_bin_url := "https://download.open-mpi.org/release/open-mpi/v5.0/openmpi-5.0.3.tar.bz2"
	openmpi_cmd_bin := exec.Command("curl", "-L", openmpi_bin_url, "-o", "binary.bin")
	err = openmpi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openmpi_src_url := "https://download.open-mpi.org/release/open-mpi/v5.0/openmpi-5.0.3.tar.bz2"
	openmpi_cmd_src := exec.Command("curl", "-L", openmpi_src_url, "-o", "source.tar.gz")
	err = openmpi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openmpi_cmd_direct := exec.Command("./binary")
	err = openmpi_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
	fmt.Println("Instalando dependencia: hwloc")
exec.Command("latte", "install", "hwloc")
	fmt.Println("Instalando dependencia: libevent")
exec.Command("latte", "install", "libevent")
	fmt.Println("Instalando dependencia: pmix")
exec.Command("latte", "install", "pmix")
}
