package main

// Gmt - Tools for manipulating and plotting geographic and Cartesian data
// Homepage: https://www.generic-mapping-tools.org/

import (
	"fmt"
	
	"os/exec"
)

func installGmt() {
	// Método 1: Descargar y extraer .tar.gz
	gmt_tar_url := "https://github.com/GenericMappingTools/gmt/releases/download/6.5.0/gmt-6.5.0-src.tar.xz"
	gmt_cmd_tar := exec.Command("curl", "-L", gmt_tar_url, "-o", "package.tar.gz")
	err := gmt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gmt_zip_url := "https://github.com/GenericMappingTools/gmt/releases/download/6.5.0/gmt-6.5.0-src.tar.xz"
	gmt_cmd_zip := exec.Command("curl", "-L", gmt_zip_url, "-o", "package.zip")
	err = gmt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gmt_bin_url := "https://github.com/GenericMappingTools/gmt/releases/download/6.5.0/gmt-6.5.0-src.tar.xz"
	gmt_cmd_bin := exec.Command("curl", "-L", gmt_bin_url, "-o", "binary.bin")
	err = gmt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gmt_src_url := "https://github.com/GenericMappingTools/gmt/releases/download/6.5.0/gmt-6.5.0-src.tar.xz"
	gmt_cmd_src := exec.Command("curl", "-L", gmt_src_url, "-o", "source.tar.gz")
	err = gmt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gmt_cmd_direct := exec.Command("./binary")
	err = gmt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: fftw")
	exec.Command("latte", "install", "fftw").Run()
	fmt.Println("Instalando dependencia: gdal")
	exec.Command("latte", "install", "gdal").Run()
	fmt.Println("Instalando dependencia: geos")
	exec.Command("latte", "install", "geos").Run()
	fmt.Println("Instalando dependencia: netcdf")
	exec.Command("latte", "install", "netcdf").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
}
