package main

// Geoip2fast - GeoIP2 country/ASN lookup tool
// Homepage: https://github.com/rabuchaim/geoip2fast

import (
	"fmt"
	
	"os/exec"
)

func installGeoip2fast() {
	// Método 1: Descargar y extraer .tar.gz
	geoip2fast_tar_url := "https://files.pythonhosted.org/packages/c6/07/2a346a6a3294e1a1f6c1852f17ae555160d6f41d8636ea00c2ae0a89a8ec/geoip2fast-1.2.2.tar.gz"
	geoip2fast_cmd_tar := exec.Command("curl", "-L", geoip2fast_tar_url, "-o", "package.tar.gz")
	err := geoip2fast_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	geoip2fast_zip_url := "https://files.pythonhosted.org/packages/c6/07/2a346a6a3294e1a1f6c1852f17ae555160d6f41d8636ea00c2ae0a89a8ec/geoip2fast-1.2.2.zip"
	geoip2fast_cmd_zip := exec.Command("curl", "-L", geoip2fast_zip_url, "-o", "package.zip")
	err = geoip2fast_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	geoip2fast_bin_url := "https://files.pythonhosted.org/packages/c6/07/2a346a6a3294e1a1f6c1852f17ae555160d6f41d8636ea00c2ae0a89a8ec/geoip2fast-1.2.2.bin"
	geoip2fast_cmd_bin := exec.Command("curl", "-L", geoip2fast_bin_url, "-o", "binary.bin")
	err = geoip2fast_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	geoip2fast_src_url := "https://files.pythonhosted.org/packages/c6/07/2a346a6a3294e1a1f6c1852f17ae555160d6f41d8636ea00c2ae0a89a8ec/geoip2fast-1.2.2.src.tar.gz"
	geoip2fast_cmd_src := exec.Command("curl", "-L", geoip2fast_src_url, "-o", "source.tar.gz")
	err = geoip2fast_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	geoip2fast_cmd_direct := exec.Command("./binary")
	err = geoip2fast_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
