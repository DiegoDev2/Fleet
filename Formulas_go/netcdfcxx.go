package main

// NetcdfCxx - C++ libraries and utilities for NetCDF
// Homepage: https://www.unidata.ucar.edu/software/netcdf/

import (
	"fmt"
	
	"os/exec"
)

func installNetcdfCxx() {
	// Método 1: Descargar y extraer .tar.gz
	netcdfcxx_tar_url := "https://github.com/Unidata/netcdf-cxx4/archive/refs/tags/v4.3.1.tar.gz"
	netcdfcxx_cmd_tar := exec.Command("curl", "-L", netcdfcxx_tar_url, "-o", "package.tar.gz")
	err := netcdfcxx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	netcdfcxx_zip_url := "https://github.com/Unidata/netcdf-cxx4/archive/refs/tags/v4.3.1.zip"
	netcdfcxx_cmd_zip := exec.Command("curl", "-L", netcdfcxx_zip_url, "-o", "package.zip")
	err = netcdfcxx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	netcdfcxx_bin_url := "https://github.com/Unidata/netcdf-cxx4/archive/refs/tags/v4.3.1.bin"
	netcdfcxx_cmd_bin := exec.Command("curl", "-L", netcdfcxx_bin_url, "-o", "binary.bin")
	err = netcdfcxx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	netcdfcxx_src_url := "https://github.com/Unidata/netcdf-cxx4/archive/refs/tags/v4.3.1.src.tar.gz"
	netcdfcxx_cmd_src := exec.Command("curl", "-L", netcdfcxx_src_url, "-o", "source.tar.gz")
	err = netcdfcxx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	netcdfcxx_cmd_direct := exec.Command("./binary")
	err = netcdfcxx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: hdf5")
exec.Command("latte", "install", "hdf5")
	fmt.Println("Instalando dependencia: netcdf")
exec.Command("latte", "install", "netcdf")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
