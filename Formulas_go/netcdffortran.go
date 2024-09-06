package main

// NetcdfFortran - Fortran libraries and utilities for NetCDF
// Homepage: https://www.unidata.ucar.edu/software/netcdf/

import (
	"fmt"
	
	"os/exec"
)

func installNetcdfFortran() {
	// Método 1: Descargar y extraer .tar.gz
	netcdffortran_tar_url := "https://github.com/Unidata/netcdf-fortran/archive/refs/tags/v4.6.1.tar.gz"
	netcdffortran_cmd_tar := exec.Command("curl", "-L", netcdffortran_tar_url, "-o", "package.tar.gz")
	err := netcdffortran_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	netcdffortran_zip_url := "https://github.com/Unidata/netcdf-fortran/archive/refs/tags/v4.6.1.zip"
	netcdffortran_cmd_zip := exec.Command("curl", "-L", netcdffortran_zip_url, "-o", "package.zip")
	err = netcdffortran_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	netcdffortran_bin_url := "https://github.com/Unidata/netcdf-fortran/archive/refs/tags/v4.6.1.bin"
	netcdffortran_cmd_bin := exec.Command("curl", "-L", netcdffortran_bin_url, "-o", "binary.bin")
	err = netcdffortran_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	netcdffortran_src_url := "https://github.com/Unidata/netcdf-fortran/archive/refs/tags/v4.6.1.src.tar.gz"
	netcdffortran_cmd_src := exec.Command("curl", "-L", netcdffortran_src_url, "-o", "source.tar.gz")
	err = netcdffortran_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	netcdffortran_cmd_direct := exec.Command("./binary")
	err = netcdffortran_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
	fmt.Println("Instalando dependencia: hdf5")
exec.Command("latte", "install", "hdf5")
	fmt.Println("Instalando dependencia: netcdf")
exec.Command("latte", "install", "netcdf")
}
