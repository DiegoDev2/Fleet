package main

// Netcdf - Libraries and data formats for array-oriented scientific data
// Homepage: https://www.unidata.ucar.edu/software/netcdf/

import (
	"fmt"
	
	"os/exec"
)

func installNetcdf() {
	// Método 1: Descargar y extraer .tar.gz
	netcdf_tar_url := "https://github.com/Unidata/netcdf-c/archive/refs/tags/v4.9.2.tar.gz"
	netcdf_cmd_tar := exec.Command("curl", "-L", netcdf_tar_url, "-o", "package.tar.gz")
	err := netcdf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	netcdf_zip_url := "https://github.com/Unidata/netcdf-c/archive/refs/tags/v4.9.2.zip"
	netcdf_cmd_zip := exec.Command("curl", "-L", netcdf_zip_url, "-o", "package.zip")
	err = netcdf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	netcdf_bin_url := "https://github.com/Unidata/netcdf-c/archive/refs/tags/v4.9.2.bin"
	netcdf_cmd_bin := exec.Command("curl", "-L", netcdf_bin_url, "-o", "binary.bin")
	err = netcdf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	netcdf_src_url := "https://github.com/Unidata/netcdf-c/archive/refs/tags/v4.9.2.src.tar.gz"
	netcdf_cmd_src := exec.Command("curl", "-L", netcdf_src_url, "-o", "source.tar.gz")
	err = netcdf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	netcdf_cmd_direct := exec.Command("./binary")
	err = netcdf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: hdf5")
	exec.Command("latte", "install", "hdf5").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
