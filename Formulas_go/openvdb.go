package main

// Openvdb - Sparse volumetric data processing toolkit
// Homepage: https://www.openvdb.org/

import (
	"fmt"
	
	"os/exec"
)

func installOpenvdb() {
	// Método 1: Descargar y extraer .tar.gz
	openvdb_tar_url := "https://github.com/AcademySoftwareFoundation/openvdb/archive/refs/tags/v11.0.0.tar.gz"
	openvdb_cmd_tar := exec.Command("curl", "-L", openvdb_tar_url, "-o", "package.tar.gz")
	err := openvdb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openvdb_zip_url := "https://github.com/AcademySoftwareFoundation/openvdb/archive/refs/tags/v11.0.0.zip"
	openvdb_cmd_zip := exec.Command("curl", "-L", openvdb_zip_url, "-o", "package.zip")
	err = openvdb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openvdb_bin_url := "https://github.com/AcademySoftwareFoundation/openvdb/archive/refs/tags/v11.0.0.bin"
	openvdb_cmd_bin := exec.Command("curl", "-L", openvdb_bin_url, "-o", "binary.bin")
	err = openvdb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openvdb_src_url := "https://github.com/AcademySoftwareFoundation/openvdb/archive/refs/tags/v11.0.0.src.tar.gz"
	openvdb_cmd_src := exec.Command("curl", "-L", openvdb_src_url, "-o", "source.tar.gz")
	err = openvdb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openvdb_cmd_direct := exec.Command("./binary")
	err = openvdb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: c-blosc")
exec.Command("latte", "install", "c-blosc")
	fmt.Println("Instalando dependencia: jemalloc")
exec.Command("latte", "install", "jemalloc")
	fmt.Println("Instalando dependencia: openexr")
exec.Command("latte", "install", "openexr")
	fmt.Println("Instalando dependencia: tbb")
exec.Command("latte", "install", "tbb")
}
