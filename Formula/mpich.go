package main

// Mpich - Implementation of the MPI Message Passing Interface standard
// Homepage: https://www.mpich.org/

import (
	"fmt"
	
	"os/exec"
)

func installMpich() {
	// Método 1: Descargar y extraer .tar.gz
	mpich_tar_url := "https://www.mpich.org/static/downloads/4.2.2/mpich-4.2.2.tar.gz"
	mpich_cmd_tar := exec.Command("curl", "-L", mpich_tar_url, "-o", "package.tar.gz")
	err := mpich_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mpich_zip_url := "https://www.mpich.org/static/downloads/4.2.2/mpich-4.2.2.zip"
	mpich_cmd_zip := exec.Command("curl", "-L", mpich_zip_url, "-o", "package.zip")
	err = mpich_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mpich_bin_url := "https://www.mpich.org/static/downloads/4.2.2/mpich-4.2.2.bin"
	mpich_cmd_bin := exec.Command("curl", "-L", mpich_bin_url, "-o", "binary.bin")
	err = mpich_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mpich_src_url := "https://www.mpich.org/static/downloads/4.2.2/mpich-4.2.2.src.tar.gz"
	mpich_cmd_src := exec.Command("curl", "-L", mpich_src_url, "-o", "source.tar.gz")
	err = mpich_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mpich_cmd_direct := exec.Command("./binary")
	err = mpich_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: hwloc")
	exec.Command("latte", "install", "hwloc").Run()
	fmt.Println("Instalando dependencia: libfabric")
	exec.Command("latte", "install", "libfabric").Run()
}
