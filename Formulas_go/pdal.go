package main

// Pdal - Point data abstraction library
// Homepage: https://www.pdal.io/

import (
	"fmt"
	
	"os/exec"
)

func installPdal() {
	// Método 1: Descargar y extraer .tar.gz
	pdal_tar_url := "https://github.com/PDAL/PDAL/releases/download/2.7.2/PDAL-2.7.2-src.tar.bz2"
	pdal_cmd_tar := exec.Command("curl", "-L", pdal_tar_url, "-o", "package.tar.gz")
	err := pdal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pdal_zip_url := "https://github.com/PDAL/PDAL/releases/download/2.7.2/PDAL-2.7.2-src.tar.bz2"
	pdal_cmd_zip := exec.Command("curl", "-L", pdal_zip_url, "-o", "package.zip")
	err = pdal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pdal_bin_url := "https://github.com/PDAL/PDAL/releases/download/2.7.2/PDAL-2.7.2-src.tar.bz2"
	pdal_cmd_bin := exec.Command("curl", "-L", pdal_bin_url, "-o", "binary.bin")
	err = pdal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pdal_src_url := "https://github.com/PDAL/PDAL/releases/download/2.7.2/PDAL-2.7.2-src.tar.bz2"
	pdal_cmd_src := exec.Command("curl", "-L", pdal_src_url, "-o", "source.tar.gz")
	err = pdal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pdal_cmd_direct := exec.Command("./binary")
	err = pdal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gdal")
exec.Command("latte", "install", "gdal")
	fmt.Println("Instalando dependencia: hdf5")
exec.Command("latte", "install", "hdf5")
	fmt.Println("Instalando dependencia: laszip")
exec.Command("latte", "install", "laszip")
	fmt.Println("Instalando dependencia: libgeotiff")
exec.Command("latte", "install", "libgeotiff")
	fmt.Println("Instalando dependencia: libpq")
exec.Command("latte", "install", "libpq")
	fmt.Println("Instalando dependencia: libxml2")
exec.Command("latte", "install", "libxml2")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: proj")
exec.Command("latte", "install", "proj")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: libunwind")
exec.Command("latte", "install", "libunwind")
}
