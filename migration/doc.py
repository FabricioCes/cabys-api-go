import csv
import json

def parse_tax_rate(value):
    value = value.strip().lower()
    if value in ["exento", "exempt", "exención"]:
        return 0.0
    try:
        return float(value)
    except ValueError:
        print(f"⚠️ Valor de impuesto inválido: '{value}', se omitirá.")
        return None

def convert_csv_to_json(csv_file, json_file):
    cabys_data = []

    with open(csv_file, mode='r', encoding='utf-8') as f:
        reader = csv.reader(f)
        for row in reader:
            if len(row) != 3:
                continue  # Omitir líneas incompletas

            tax = parse_tax_rate(row[2])
            if tax is None:
                continue  # Omitir si no es convertible

            cabys_data.append({
                "id": row[0].strip(),
                "description": row[1].strip(),
                "tax_percent": tax
            })

    with open(json_file, mode='w', encoding='utf-8') as f:
        json.dump(cabys_data, f, indent=2, ensure_ascii=False)

    print(f"✅ Generado {json_file} con {len(cabys_data)} códigos.")

if __name__ == "__main__":
    convert_csv_to_json("cabys.csv", "cabys.json")
