package main

// Hdf5AT18 - File format designed to store large amounts of data
// Homepage: https://www.hdfgroup.org/HDF5

import (
	"fmt"
	
	"os/exec"
)

func installHdf5AT18() {
	// Método 1: Descargar y extraer .tar.gz
	hdf5at18_tar_url := "https://support.hdfgroup.org/ftp/HDF5/releases/hdf5-1.8/hdf5-1.8.23/src/hdf5-1.8.23.tar.bz2"
	hdf5at18_cmd_tar := exec.Command("curl", "-L", hdf5at18_tar_url, "-o", "package.tar.gz")
	err := hdf5at18_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hdf5at18_zip_url := "https://support.hdfgroup.org/ftp/HDF5/releases/hdf5-1.8/hdf5-1.8.23/src/hdf5-1.8.23.tar.bz2"
	hdf5at18_cmd_zip := exec.Command("curl", "-L", hdf5at18_zip_url, "-o", "package.zip")
	err = hdf5at18_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hdf5at18_bin_url := "https://support.hdfgroup.org/ftp/HDF5/releases/hdf5-1.8/hdf5-1.8.23/src/hdf5-1.8.23.tar.bz2"
	hdf5at18_cmd_bin := exec.Command("curl", "-L", hdf5at18_bin_url, "-o", "binary.bin")
	err = hdf5at18_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hdf5at18_src_url := "https://support.hdfgroup.org/ftp/HDF5/releases/hdf5-1.8/hdf5-1.8.23/src/hdf5-1.8.23.tar.bz2"
	hdf5at18_cmd_src := exec.Command("curl", "-L", hdf5at18_src_url, "-o", "source.tar.gz")
	err = hdf5at18_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hdf5at18_cmd_direct := exec.Command("./binary")
	err = hdf5at18_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libaec")
	exec.Command("latte", "install", "libaec").Run()
}
