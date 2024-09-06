package main

// Pgrouting - Provides geospatial routing for PostGIS/PostgreSQL database
// Homepage: https://pgrouting.org/

import (
	"fmt"
	
	"os/exec"
)

func installPgrouting() {
	// Método 1: Descargar y extraer .tar.gz
	pgrouting_tar_url := "https://github.com/pgRouting/pgrouting/releases/download/v3.6.2/pgrouting-3.6.2.tar.gz"
	pgrouting_cmd_tar := exec.Command("curl", "-L", pgrouting_tar_url, "-o", "package.tar.gz")
	err := pgrouting_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pgrouting_zip_url := "https://github.com/pgRouting/pgrouting/releases/download/v3.6.2/pgrouting-3.6.2.zip"
	pgrouting_cmd_zip := exec.Command("curl", "-L", pgrouting_zip_url, "-o", "package.zip")
	err = pgrouting_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pgrouting_bin_url := "https://github.com/pgRouting/pgrouting/releases/download/v3.6.2/pgrouting-3.6.2.bin"
	pgrouting_cmd_bin := exec.Command("curl", "-L", pgrouting_bin_url, "-o", "binary.bin")
	err = pgrouting_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pgrouting_src_url := "https://github.com/pgRouting/pgrouting/releases/download/v3.6.2/pgrouting-3.6.2.src.tar.gz"
	pgrouting_cmd_src := exec.Command("curl", "-L", pgrouting_src_url, "-o", "source.tar.gz")
	err = pgrouting_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pgrouting_cmd_direct := exec.Command("./binary")
	err = pgrouting_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: cgal")
	exec.Command("latte", "install", "cgal").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: postgis")
	exec.Command("latte", "install", "postgis").Run()
	fmt.Println("Instalando dependencia: postgresql@14")
	exec.Command("latte", "install", "postgresql@14").Run()
}
